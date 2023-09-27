import { UUID } from "crypto";

class User {
    id: UUID;
    username: string;
    email: string;
    phone: string;
    birthday:Date;
    name: string;
    lastName: string;
    patronimyc: string;
    createdAt: Date;
    updatedAt: Date;
    password: string;

    constructor(id: UUID, username: string, email: string, phone:string, birthday:Date, name: string, lastName: string, patronimyc: string, createdAt: Date, updatedAt: Date) {
        this.id = id;
        this.username = username;
        this.email = email;
        this.phone=phone;
        this.birthday=birthday;
        this.name = name;
        this.lastName = lastName;
        this.patronimyc = patronimyc;
        this.createdAt = createdAt;
        this.updatedAt = updatedAt;
        this.password=""
    }
}
