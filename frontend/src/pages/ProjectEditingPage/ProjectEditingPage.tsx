// react imports
import React, {
    useState,
} from "react";

// component imports
import ProjectEditingMenu from "./ProjectEditingMenu/ProjectEditingMenu";
import ProjectEditingView from "./ProjectEditingView/ProjectEditingView";

// interface imports
import { Project } from "../../interfaces/Project";

// materialui imports
import { Box } from "@mui/material";

const ProjectEditingPage: React.FC = () => {
    const [project, setProject] = useState<Project>()

    return (
        <Box>
            <ProjectEditingMenu />
            <ProjectEditingView />
        </Box>
    )
}

export default ProjectEditingPage;