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
import AffinityInfo from './AffinityInfo';

const WIDTH = 560;

function createHandler(row) {
    return [
        CommonHandler.createDelete("deployment", row.namespace, row.name),
        CommonHandler.createEdit("deployment", row.namespace, row.name),
        CommonHandler.createScale('deployment',row.namespace,row.name),
        CommonHandler.createRestart('deployment',row.namespace,row.name),
        CommonHandler.createPause(row.namespace,row.name),
        CommonHandler.createResume(row.namespace,row.name),      
    ]
}

function DeploymentInfo(props) {
    const {deployment} = props;
    return (
        <Box
        sx={{ maxWidth:WIDTH}}
    >    
    <DetailLine name="Created">
        <DetailTypography>{deployment.created}</DetailTypography>
    </DetailLine>
    <DetailLine name="Name">
        <DetailTypography>{deployment.name}</DetailTypography>
    </DetailLine>
    <DetailLine name="Namespace">
        <LinkTypography>{deployment.namespace}</LinkTypography>
    </DetailLine>
    <DetailLine name="Annotations">
    <ArrayTypography items={deployment.annotations}></ArrayTypography>
    </DetailLine>
    <DetailLine name="Replicas">
        <DetailTypography>{deployment.replicas}</DetailTypography>
    </DetailLine>
    <DetailLine name="Selector">
    <ArrayTypography items={deployment.selector}></ArrayTypography>
    </DetailLine>
    <DetailLine name="StrategyType">
        <DetailTypography>{deployment.strategyType}</DetailTypography>
    </DetailLine>
    </Box>
    )
}

function PodInfo(props) {
    const columns = ["name","ready","status"];
    const {pods} = props;
    return (
        <Box sx={{ maxWidth:WIDTH}}>
    <TableContainer >
          <Table  size="small">
            <TableHead>
              <TableRow>
                { columns.map((column,index) => (
                    <TableCell key={index}>{column}</TableCell>
                ))}
               
              </TableRow>
            </TableHead>
            <TableBody>
              {pods.map((row,index) => (
                <TableRow
                  key={index}
                  sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                  {
                    columns.map((column,index) => (
                        <TableCell key={index} >{row[column]}</TableCell>
                    ))
                  }
                   
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
        </Box>
    )
}

export default function DeploymentDetail(props) {
    const defaultData = {
        name:"",
        namespace:"",
        created:"",
        annotations:[],
        replicas:0,
        selector:[],
        strategyType:"",
        pods:[],
    };
    const {namespace,name,open,onClose} = props;
    const [data, setData] = React.useState(defaultData);

    React.useEffect(() => {
        async function fetchData() {
            const result = await http.deploymentDetail(namespace,name)
            setData(result);
        }
        if (open) {
            fetchData();
        }
    }, [open]);
    
    const info = () => (
        <>
             <DetailBar
                title={"Deployment:" + data.name} 
                moreHandles={createHandler(data)}/>
             <Toolbar />
             <DeploymentInfo deployment={data}/> 
             <TitleBar title="Pods" />
             <PodInfo pods={data.pods}/>
             <TitleBar title="Affinity" />
             <AffinityInfo />
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