import * as React from 'react';
import DetailTypography from "./DetailTypography";
import { Grid ,Divider} from '@mui/material';

export default function DetailLine(props) {
    return (
        <>
            <Grid container spacing={2} minHeight={40}>
                <Grid item xs={3}>
                    <DetailTypography>{props.name}</DetailTypography>
                </Grid>
                <Grid item xs={9}>
                    {props.children}
                </Grid>
            </Grid>
            <Divider />
        </>
    )
}