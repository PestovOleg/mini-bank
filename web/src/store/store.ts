import { AccountStore } from "./account.store";
import { UserStore } from "./user.store";
import { ToggleStore } from "./toggle.store";

export class Store{
    userStore: UserStore;
    accountStore:AccountStore;
    toggleStore:ToggleStore;

    constructor(){
        this.userStore=new UserStore(this);
        this.accountStore=new AccountStore(this);
        this.toggleStore=new ToggleStore(this);
    }
}

const store = new Store();
export default store;