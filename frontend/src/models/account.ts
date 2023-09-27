import { UUID } from "crypto";

export class Account{
    id: UUID | null;
    account: string;
    amount: number;
    currency: string;
    name: string;
    createdAt: Date;
    interestRate: number;
    userID: UUID | null;

    constructor(id:UUID,account:string,amount:number,currency:string,name:string,createdAt:Date,interestRate:number,userID:UUID){
        this.id=id;
        this.account=account;
        this.amount=amount;
        this.currency=currency;
        this.name=name;
        this.createdAt=createdAt;
        this.interestRate=interestRate;
        this.userID=userID;
    }
}