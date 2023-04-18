import * as React from 'react';
import http from '../../http/detail';
import { 
    Drawer,
    Box, 
    Toolbar,
    Table,
    TableContainer,
    TableHead,
    TableRow,
    TableCell,
    TableBody,
    Paper,
} from "@mui/material";
import DetailTypography from "../ext/DetailTypography";
import LinkTypography from "../ext/LinkTypography";
import DetailLine from "../ext/DetailLine";
import ArrayTypography from "../ext/ArrayTypography";
import TitleBar from "../ext/TitleBar";
import DetailBar from "../ext/DetailBar";
import CommonHandler from "../../page/common";
import EventInfo from './EventInfo';
import ConfigReader from '../ConfigReader';
const WIDTH = 560;
function createHandler(row) {
    return [
        CommonHandler.createDelete("configmap", row.namespace, row.name),
        CommonHandler.createEdit("configmap", row.namespace, row.name),         
    ]
}

function ConfigmapInfo(props) {
    const {configmap} = props;
    return (
        <Box
        sx={{ maxWidth:WIDTH}}
    >    
    <DetailLine name="Created">
        <DetailTypography>{configmap.created}</DetailTypography>
    </DetailLine>
    <DetailLine name="Name">
        <DetailTypography>{configmap.name}</DetailTypography>
    </DetailLine>
    <DetailLine name="Namespace">
        <LinkTypography>{configmap.namespace}</LinkTypography>
    </DetailLine>
    <DetailLine name="Annotations">
    <ArrayTypography items={configmap.annotations}></ArrayTypography>
    </DetailLine>
    </Box>
    )
}

function DataInfo(props){
    const {data} = props;
    const keys = Object.keys(data);
    
    return (
        <>
         {
            
            keys.map((key,index) => (
                <>
                <DetailTypography>{key}</DetailTypography>
                <ConfigReader
                    code={data[key]}
                />
                </>
            ))
               
         }
         
        </>
    )
}

export default function ConfigmapDetail(props) {
    const defaultData = {
        name:"",
        namespace:"",
        created:"",
        annotations:[],
        data:{},
    };
    const {namespace,name,open,onClose} = props;
    const [data, setData] = React.useState(defaultData);

    React.useEffect(() => {
        async function fetchData() {
            const result = await http.configmapDetail(namespace,name)
            setData(result);
        }
        if (open) {
            fetchData();
        }
    }, [open]);
    
    const info = () => (
        <>
           
             <DetailBar
                title={"Configmap:" + data.name} 
                moreHandles={createHandler(data)}/>
             <Toolbar />
             <ConfigmapInfo configmap={data}/>
             <TitleBar title="Data" />
             <DataInfo data={data.data}/> 
             <TitleBar title="Events" />
             <EventInfo />
             
         </>
     ) 

    return (
        <>
            <Drawer
                anchor="right"
                open={open}
                onClose={onClose}
            >
                {info()}
            </Drawer>
        </>
    )
}