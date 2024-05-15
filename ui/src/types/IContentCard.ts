import { IComment } from "./IComment";

export interface IContentCard {
  id: string;
  imageUri: string;
  title: string;
  subtitle: string;
  body: string;
  author: string;
  priority: number;
  publishDate: string;
  comments: IComment[];
}
