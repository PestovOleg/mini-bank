import { IUser } from "../models/types";
import { makeAutoObservable } from "mobx";
import { runInAction } from "mobx";
import { EMPTY_USER } from "../const/empties";
import { Store } from "./store";

const URL = process.env.REACT_APP_URL;
//const URL = "http://localhost/api/v1"

export class UserStore {
    public User: IUser;

    public isAuth: boolean = false;

    public authError: string = "";

    public signUpSuccess: boolean = false;

    constructor(private rootStore: Store) {
        makeAutoObservable(this);
        this.User = { ...EMPTY_USER };
        this.isAuth = false;
    }

    public async login(username: string, password: string): Promise<void> {
        const base64Credentials = btoa(username + ":" + password);
        console.log("Вывод из STORE: username:", username, " password: ", password);
        try {
            const response = await fetch(`${URL}/auth/login`, {
                method: "POST",
                headers: {
                    Authorization: "Basic " + base64Credentials,
                },
            });

            if (!response.ok) {
                this.authError = "Failed to login.";
                throw new Error(response.statusText);
            }

            const res: any = await response.json();

            runInAction(() => {
                this.isAuth = true;
                this.User.id = res.id;
                this.User.username = username;
                this.User.password = password;
                this.User.token = "Basic " + base64Credentials;
            });
        } catch (error) {
            console.error("Login error:", error);
            // Handle or throw the error as needed
        }
    }

    public logout(): void {
        this.isAuth = false;
        this.User = { ...EMPTY_USER };
        console.log("logout")
    }

    public async signup(
        firstName: string,
        lastName: string,
        patronymic: string,
        email: string,
        username: string,
        password: string,
        phone: string,
        birthday: string
    ): Promise<void> {
        const userData = {
            email,
            last_name: lastName,
            name: firstName,
            password,
            patronymic,
            username,
            phone: phone.replace(/\D/g, ""),
            birthday: birthday,
        };

        try {
            const response = await fetch(`${URL}/mgmt`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(userData),
            });

            if (!response.ok) {
                throw new Error(response.statusText);
            }

            const res: any = await response.json();
            runInAction(() => {
                this.signUpSuccess = true;
            });
        } catch (error) {
            console.error("There was a problem with the fetch operation:", error);
        }
    }

    public async getUser(): Promise<void> {

        try {
            const response = await fetch(`${URL}/users/${this.User.id}`, {
                headers: {
                    Authorization: this.User.token,
                    'Cache-Control': 'no-cache',
                },
            });

            if (!response.ok) {
                throw new Error(response.statusText);
            }

            const res: any = await response.json();

            runInAction(() => {
                this.User.username = res.username;
                this.User.email = res.email;
                this.User.phone = res.phone;
                this.User.birthday = res.birthday;
                this.User.name = res.name;
                this.User.lastName = res.last_name;
                this.User.patronimyc = res.patronymic;
                this.User.createdAt = new Date(res.created_at);
            });
        } catch (error) {
            console.log(this.User.id);
            console.error("There was a problem with the fetch operation:", error);
        }
    }

    public async updateUser(email: string, phone: string): Promise<void> {

        const userData = {
            email,
            phone,
        };

        try {
            const response = await fetch(`${URL}/users/${this.User.id}`, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: this.User.token,
                },
                body: JSON.stringify(userData),
            });

            if (!response.ok) {
                throw new Error(response.statusText);
            }

            const res: any = await response.json();
            runInAction(() => {
                this.getUser();
                
            });
        } catch (error) {
            console.error("There was a problem with the fetch operation:", error);
        }
    }

    public async deleteUser(): Promise<void> {        

        try {
            const response = await fetch(`${URL}/mgmt/${this.User.id}`, {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: this.User.token,
                },
            });

            if (!response.ok) {
                throw new Error(response.statusText);
            }

            const res: any = await response.json();
            runInAction(() => {
                this.logout();
                
            });
        } catch (error) {
            console.error("There was a problem with the fetch operation:", error);
        }
    }
}
