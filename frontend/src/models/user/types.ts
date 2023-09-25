export type UUID = string;

export interface IUser {
    id: UUID|null
    username: string
    email: string 
    name: string 
    lastName: string 
    patronimyc: string 
    createdAt: Date
    password:string
} 