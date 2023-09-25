import { UserStore } from "./user.store";

class Store{
    userStore: UserStore;

    constructor(){
        this.userStore=new UserStore();
    }
}

const store = new Store();
export default store;