export type UUID = string;

export interface IUser {
    id: UUID | null;
    username: string;
    email: string;
    phone: string;
    birthday: Date;
    name: string;
    lastName: string;
    patronimyc: string;
    createdAt: Date;
    password: string;
    token: string;
}

export interface IAccount {
    id: UUID | null;
    account: string;
    amount: number;
    currency: string;
    name: string;
    createdAt: Date;
    interestRate: number;
    userID: UUID | null;
}

export interface IFeatureEnvironment  {
    name: string;
    enabled: boolean;
    type: string;
  };
  
export  interface IFeature  {
    name: string;
    description: string;
    environments: IFeatureEnvironment[];
  };
  
