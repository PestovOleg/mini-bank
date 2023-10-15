
export class Feature {
    name: string;
    description: string;
    environments: FeatureEnvironment[];

    constructor(name:string,description:string,environments:FeatureEnvironment[]){
        this.name=name;
        this.description=description;
        this.environments=environments;
    }
}

export class FeatureEnvironment {
    name: string;
    enabled: boolean;
    type: string;

    constructor(name:string,enabled:boolean,type:string){
        this.name=name;
        this.enabled=enabled;
        this.type=type;
    }
}