import { IAccount, IUser } from "../models/types";
import { action, makeAutoObservable, observable } from "mobx";
import { runInAction } from "mobx";
import { EMPTY_USER,EMPTY_ACCOUNT } from "../const/empties"
import { Store } from './store';

const URL = "http://localhost/api/v1";

export class AccountStore {
    public Accounts: IAccount[];

    constructor(private rootStore: Store) {
        makeAutoObservable(this);
        this.Accounts = [{ ...EMPTY_ACCOUNT }];
    }
    
    public async getList(): Promise<void> {
        const base64Credentials = btoa(this.rootStore.userStore.User.username + ':' + this.rootStore.userStore.User.password);
        
        try {
            const response = await fetch(`${URL}/users/${this.rootStore.userStore.User.id}/accounts`, {
                headers: {
                    'Authorization': 'Basic ' + base64Credentials
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
    
}
