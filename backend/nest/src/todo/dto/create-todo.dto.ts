import { IsBoolean, IsString } from 'class-validator';

export class CreateTodoDto {
  @IsString()
  title;

  @IsString()
  note;

  @IsBoolean()
  completed;
}
