import { AWSServiceEnums } from "../enums/awsEnums";

export interface Project {
    id: string;
    name: string;
    owner_id: string;
}

export interface ProjectViewAWSComponent {
    id: string;
    type: AWSServiceEnums;
    img: string;
    owner_id: string;
    project_id: string;
    dimensions: {
        x: number;
        y: number;
        height: number;
        width: number;
    };
};
