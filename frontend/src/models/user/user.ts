import { UUID } from "crypto";

class User {
    id: UUID;
    username: string;
    email: string
    name: string
    lastName: string
    patronimyc: string
    createdAt: Date
    updatedAt: Date
    password: string

    constructor(id: UUID, username: string, email: string, name: string, lastName: string, patronimyc: string, createdAt: Date, updatedAt: Date) {
        this.id = id;
        this.username = username;
        this.email = email;
        this.name = name;
        this.lastName = name;
        this.patronimyc = patronimyc;
        this.createdAt = createdAt;
        this.updatedAt = updatedAt;
        this.password=""
    }
}
