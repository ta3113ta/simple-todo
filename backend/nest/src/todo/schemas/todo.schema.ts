import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';
import { v4 as uuidv4 } from 'uuid';

export type TodoDocument = Todo & Document;

@Schema()
export class Todo {
  @Prop({ default: uuidv4 })
  id: string;

  @Prop({ required: true })
  title: string;

  @Prop()
  note: string;

  @Prop({ default: false })
  completed: boolean;
}

export const TodoSchema = SchemaFactory.createForClass(Todo);
