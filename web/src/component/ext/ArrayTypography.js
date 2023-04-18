import { Margin } from "@mui/icons-material";
import { 
    Chip 
} from "@mui/material";
export default function ArrayTypography(props) {
    return (
        <>
        {
            props.items.map((item,index) => (
                <Chip key={index} label={item} color="default" variant="outlined" 
                    sx={{marginTop:"5px",marginBottom:"5px"}}/>
            ))
        }
        </>
    )
}