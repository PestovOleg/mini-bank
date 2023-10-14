import { IFeature } from "../models/types";
import {  makeAutoObservable } from "mobx";
import { runInAction } from "mobx";
import { EMPTY_FEATURE } from "../const/empties";
import { Store } from "./store";
import { Feature, FeatureEnvironment } from "../models/feature";

const URL = process.env.REACT_APP_URL;
//const URL = "http://localhost/api/v1"

export class ToggleStore {
    public Toggles: IFeature[];
    public Environment:string|undefined;

    constructor(private rootStore: Store) {
        makeAutoObservable(this);
        this.Toggles=[EMPTY_FEATURE]
        this.Environment=process.env.NODE_ENV
    }

    
    public async getToggles(): Promise<void> {

        try {
            const response = await fetch(`${URL}/uproxy`, {               
            });

            if (!response.ok) {
                throw new Error(response.statusText);
            }

            const res: any = await response.json();

            runInAction(() => {
                this.Toggles = res.features.map((feature: any) => {
                    return new Feature(
                        feature.name,
                        feature.description,
                        feature.environments.map((env: any) => new FeatureEnvironment(env.name, env.enabled, env.type))
                    );
                });
                console.log("Тоглы "+this.Toggles)
                console.log(this.Toggles)
            });
        } catch (error) {
            console.error("There was a problem with the fetch operation:", error);
        }
    }

    public getFeature(name: string): boolean | null {
        const feature = this.Toggles.find((feature) => feature.name === name);
        if (feature) {
            const env = feature.environments.find((env) => env.name === this.Environment);
            if (env) {
                return env.enabled;
            }
        }
        return null;
    }
    
    
}