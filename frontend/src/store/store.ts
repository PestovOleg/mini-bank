import { AccountStore } from "./account.store";
import { UserStore } from "./user.store";

export class Store{
    userStore: UserStore;
    accountStore:AccountStore;

    constructor(){
        this.userStore=new UserStore(this);
        this.accountStore=new AccountStore(this);
    }
}

const store = new Store();
export default store;