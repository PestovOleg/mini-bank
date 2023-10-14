import { IAccount } from "../models/types";
import {  makeAutoObservable } from "mobx";
import { runInAction } from "mobx";
import { EMPTY_ACCOUNT } from "../const/empties"
import { Store } from './store';

const URL = process.env.REACT_APP_URL;
//const URL = "http://localhost/api/v1";

export class AccountStore {
    public Accounts: IAccount[];

    constructor(private rootStore: Store) {
        makeAutoObservable(this);
        this.Accounts = [{ ...EMPTY_ACCOUNT }];
    }

    public async getList(): Promise<void> {

        try {
            const response = await fetch(`${URL}/users/${this.rootStore.userStore.User.id}/accounts`, {
                headers: {
                    'Authorization': this.rootStore.userStore.User.token,
                }
            });

            if (!response.ok) {
                throw new Error('Failed to fetch list of accounts');
            }

            const res: any[] = await response.json();

            runInAction(() => {
                this.Accounts = [];
                res.forEach((item: any) => {
                    const newAccount: IAccount = {
                        id: item.id,
                        account: item.account,
                        amount: item.amount,
                        currency: item.currency,
                        name: item.name,
                        createdAt: new Date(item.created_at),
                        interestRate: item.interest_rate,
                        userID: item.user_id
                    };
                    this.Accounts.push(newAccount);
                });
                console.log(this.Accounts)
            });
        } catch (error) {
            console.error('There was a problem with the fetch operation:', error);
        }
    }

    public async getAccountInfo(id: string): Promise<void> {

        try {
            const response = await fetch(`${URL}/users/${this.rootStore.userStore.User.id}/accounts/${id}`, {
                headers: {
                    'Authorization': this.rootStore.userStore.User.token,
                }
            });

            if (!response.ok) {
                throw new Error('Failed to get account info');
            }

            const res: any = await response.json();

            runInAction(() => {
                const accountToUpdate = this.Accounts.find(account => account.id === id);

                if (accountToUpdate) {
                    accountToUpdate.amount = res.amount;
                    accountToUpdate.currency = res.currency;
                    accountToUpdate.createdAt = res.created_at;
                    accountToUpdate.interestRate = res.interest_rate;
                    accountToUpdate.account = res.account;
                    accountToUpdate.userID = res.user_id;
                    accountToUpdate.name = res.name;
                }
                console.log(this.Accounts)
            });
        } catch (error) {
            console.error('There was a problem with the fetch operation:', error);
        }
    }

    public async openAccount(
        currency: string,
        accountName: string
    ): Promise<void> {
        const accountData = {
            currency: String(currency),
            name: accountName,
        };

        try {
            const response = await fetch(`${URL}/users/${this.rootStore.userStore.User.id}/accounts`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    'Authorization': this.rootStore.userStore.User.token,
                },
                body: JSON.stringify(accountData),
            });

            if (!response.ok) {
                throw new Error(response.statusText);
            }

            const res: any = await response.json();
            runInAction(() => {
                this.getList();
            });
        } catch (error) {
            console.error("There was a problem with the fetch operation:", error);
        }
    }

    public async updateAccount(
        accountName: string,
        id:string
    ): Promise<void> {
        const accountData = {
            name: accountName,
            id
        };

        try {
            const response = await fetch(`${URL}/users/${this.rootStore.userStore.User.id}/accounts/${id}`, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                    'Authorization': this.rootStore.userStore.User.token,
                },
                body: JSON.stringify(accountData),
            });

            if (!response.ok) {
                throw new Error(response.statusText);
            }

            runInAction(() => {
                this.getAccountInfo(id);
            });
        } catch (error) {
            console.error("There was a problem with the fetch operation:", error);
        }
    }

    public async closeAccount(id:string): Promise<void> {
       

        try {
            const response = await fetch(`${URL}/users/${this.rootStore.userStore.User.id}/accounts/${id}`, {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json",
                    'Authorization': this.rootStore.userStore.User.token,
                },
            });

            if (!response.ok) {
                throw new Error(response.statusText);
            }

            runInAction(() => {
                this.getList();
            });
        } catch (error) {
            console.error("There was a problem with the fetch operation:", error);
        }
    }

    public async topUpAccount(
        amount: number,
        id: string
    ): Promise<void> {
        const accountData = {
            amount,
            id
        };

        try {
            const response = await fetch(`${URL}/users/${this.rootStore.userStore.User.id}/accounts/${id}/topup`, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                    'Authorization': this.rootStore.userStore.User.token,
                },
                body: JSON.stringify(accountData),
            });

            if (!response.ok) {
                throw new Error(response.statusText);
            }

            const res: any = await response.json();
            runInAction(() => {
                this.getAccountInfo(id);
            });
        } catch (error) {
            console.error("There was a problem with the fetch operation:", error);
        }
    }

    public async withdrawAccount(
        amount: number,
        id: string
    ): Promise<void> {
        const accountData = {
            amount,
            id
        };

        try {
            const response = await fetch(`${URL}/users/${this.rootStore.userStore.User.id}/accounts/${id}/withdraw`, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                    'Authorization': this.rootStore.userStore.User.token,
                },
                body: JSON.stringify(accountData),
            });

            if (!response.ok) {
                throw new Error(response.statusText);
            }

            const res: any = await response.json();
            runInAction(() => {
                this.getAccountInfo(id);
            });
        } catch (error) {
            console.error("There was a problem with the fetch operation:", error);
        }
    }

}
