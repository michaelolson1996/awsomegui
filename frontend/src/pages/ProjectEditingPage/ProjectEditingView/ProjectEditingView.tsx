import React, { 
    useEffect,
    useState,
} from 'react';

// interface imports
import { 
    Project, 
    ProjectViewAWSComponent 
} from '../../../interfaces/Project';

import { AWSServiceEnums } from '../../../enums/awsEnums';

// mui imports
import Box from '@mui/material/Box';

const ProjectEditingView: React.FC = (props) => {
    const [components, setComponents] = useState<ProjectViewAWSComponent[]>([]);

    const hello: ProjectViewAWSComponent = {
        id: "",
        type: AWSServiceEnums['VPC'],
        img: AWSServiceEnums.VPC,
        owner_id: AWSServiceEnums.EC2.toString(),
        project_id: AWSServiceEnums['EC2'].toString(),
        dimensions: {
            x: 0,
            y: 0,
            height: 0,
            width: 0
        }
    }

    useEffect(() => {

    }, []);
    console.log(hello)
    return (
        <Box
            sx={{
                height: "",
            }}
        >
            {/* {console.log(hello)} */}
        </Box>
    )
}

export default ProjectEditingView