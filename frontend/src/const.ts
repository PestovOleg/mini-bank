import { IUser } from "./models/user/types";

export const EMPTY_USER: IUser = {
    id: '',
    username: '',
    email: '',
    name: '',
    lastName: '',
    patronimyc: '',
    createdAt: new Date(0),
    password:'',
};