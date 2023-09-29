import { IAccount, IUser } from "../models/types";

export const EMPTY_USER: IUser = {
    id: '',
    username: '',
    email: '',
    phone: '',
    birthday: new Date(0),
    name: '',
    lastName: '',
    patronimyc: '',
    createdAt: new Date(0),
    password:'',
};

export const EMPTY_ACCOUNT: IAccount={
    id: '',
    account: '',
    amount: 0,
    currency: '',
    name: '',
    createdAt: new Date(0),
    interestRate: 0,
    userID: '',
}