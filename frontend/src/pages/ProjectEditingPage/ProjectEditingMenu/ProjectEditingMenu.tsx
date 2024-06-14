
// materialui imports
import { 
    IconButton 
} from "@mui/material"

import AddCircleIcon from "@mui/icons-material/AddCircle"

const ProjectEditingMenu: React.FC = () => {
    return (
        <IconButton>
            <AddCircleIcon
                sx = {{
                    height: "4vh",
                    width: "4vh"
                }}
            />
        </IconButton>
    )
}

export default ProjectEditingMenu
