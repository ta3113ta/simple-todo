import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { CreateTodoDto } from './dto/create-todo.dto';
import { UpdateTodoDto } from './dto/update-todo.dto';
import { TodoDocument, Todo } from './schemas/todo.schema';

@Injectable()
export class TodoService {
  constructor(@InjectModel(Todo.name) private todoModel: Model<TodoDocument>) {}

  async create(createTodoDto: CreateTodoDto): Promise<Todo> {
    const createTodo = new this.todoModel(createTodoDto);
    return createTodo.save();
  }

  async findAll(): Promise<Todo[]> {
    const todos = await this.todoModel.find().exec();
    return todos;
  }

  async findOne(id: string) {
    const todo = await this.todoModel.findOne({ id });
    return todo;
  }

  async update(id: string, updateTodoDto: UpdateTodoDto) {
    const todo = await this.todoModel.findOneAndUpdate({ id }, updateTodoDto);
    return todo;
  }

  async remove(id: string) {
    const todo = await this.todoModel.findOneAndRemove({ id });
    if (!todo) {
      throw new NotFoundException('No todo id found');
    }
    return todo;
  }
}
