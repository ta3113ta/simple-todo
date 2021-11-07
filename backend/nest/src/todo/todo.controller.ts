import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
  ParseUUIDPipe,
} from '@nestjs/common';
import { TodoService } from './todo.service';
import { CreateTodoDto } from './dto/create-todo.dto';
import { UpdateTodoDto } from './dto/update-todo.dto';

@Controller('todo')
export class TodoController {
  constructor(private readonly todoService: TodoService) {}

  @Post()
  async create(@Body() createTodoDto: CreateTodoDto) {
    const todo = await this.todoService.create(createTodoDto);
    return { success: true, data: todo };
  }

  @Get()
  async findAll() {
    const todos = await this.todoService.findAll();
    return { success: true, data: todos };
  }

  @Get(':id')
  async findOne(@Param('id', new ParseUUIDPipe()) id: string) {
    const todo = await this.todoService.findOne(id);
    return { success: true, data: todo };
  }

  @Patch(':id')
  async update(
    @Param('id', new ParseUUIDPipe()) id: string,
    @Body() updateTodoDto: UpdateTodoDto,
  ) {
    const todo = await this.todoService.update(id, updateTodoDto);
    return { success: true, data: todo };
  }

  @Delete(':id')
  async remove(@Param('id', new ParseUUIDPipe()) id: string) {
    const todo = await this.todoService.remove(id);
    return { success: true, data: todo };
  }
}
