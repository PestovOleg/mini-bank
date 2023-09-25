import { IUser } from "../models/user/types";
import { action, makeAutoObservable, observable } from "mobx";
import { runInAction } from "mobx";
import { EMPTY_USER } from "../const"

const URL = "http://localhost/api/v1";

export class UserStore {
    public User: IUser;

    public isAuth: boolean = false;

    public authError: string = "";

    public signUpSuccess: boolean=false

    constructor() {
        makeAutoObservable(this);
        this.User = { ...EMPTY_USER };
        this.isAuth = false
    }

    public async login(username: string, password: string): Promise<void> {
        const base64Credentials = btoa(username + ':' + password);
        console.log("Вывод из STORE: username:", username, " password: ", password)
        try {
            const response = await fetch(`${URL}/users`, {
                headers: {
                    'Authorization': 'Basic ' + base64Credentials
                }
            });

            if (!response.ok) {
                this.authError = "Failed to login."
                throw new Error("Failed to login.");
            }

            const res: string = await response.json();

            runInAction(() => {
                this.isAuth = true;
                this.User.id = res;
                this.User.username = username;
                this.User.password = password;
            });

        } catch (error) {
            console.error("Login error:", error);
            // Handle or throw the error as needed
        }
    }

    public logout(): void {
        this.isAuth = false;
        this.User = { ...EMPTY_USER }
    }

    public async signup(firstName: string, lastName: string, patronymic: string, email: string, username: string, password: string): Promise<void> {
        const userData = {
            email,
            lastName,
            name: firstName,
            password,
            patronymic,
            username
        };

        try {
            const response = await fetch(`${URL}/users`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(userData)
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }

            const res: any = await response.json();
            runInAction(() => {
                this.User = res;
                this.signUpSuccess=true;
            });

        } catch (error) {
            console.error('There was a problem with the fetch operation:', error);
        }
    }


    public getUser(): void {
        const base64Credentials = btoa(this.User.username + ':' + this.User.password);

        fetch(`${URL}/users/35ec1e95-87d3-42fe-8158-bf59c78b9e26`, {
            headers: {
                'Authorization': 'Basic ' + base64Credentials
            }
        }).then(response => response.json())
            .then((res: any) => {
                runInAction(() => {
                    this.User = res;
                    console.log(this.User)
                });
            });
    }
}

const userStore = new UserStore();
export default userStore;