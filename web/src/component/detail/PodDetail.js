import { 
    Drawer,
    Box, 
    Toolbar
} from "@mui/material";
import * as React from "react";
import DetailTypography from "../ext/DetailTypography";
import LinkTypography from "../ext/LinkTypography";
import DetailLine from "../ext/DetailLine";
import StatusTypography from "../ext/StatusTypography";
import ArrayTypography from "../ext/ArrayTypography";
import TitleBar from "../ext/TitleBar";
import DetailBar from "../ext/DetailBar";
import CommonHandler from "../../page/common";
import http from '../../http/detail';
import EventInfo from "./EventInfo";
import AffinityInfo from "./AffinityInfo";

const WIDTH  = 560;

function PodInfo(props) {
    const {pod} = props;
    return (
        <Box
        sx={{ width:WIDTH,padding:"0px 10px"}}
    >       
             <DetailLine name="Created">
                <DetailTypography>{pod.created}</DetailTypography>
            </DetailLine>
            <DetailLine name="Name">
                <DetailTypography>{pod.name}</DetailTypography>
            </DetailLine>
            <DetailLine name="Namespace">
                <LinkTypography>{pod.namespace}</LinkTypography>
            </DetailLine>
            <DetailLine name="Labels">
                <ArrayTypography items={pod.labels}></ArrayTypography>
            </DetailLine>
            <DetailLine name="Controlled By">
                <DetailTypography>{pod.controlledBy.kind}</DetailTypography>
                <LinkTypography>{pod.controlledBy.name}</LinkTypography>
            </DetailLine>
            <DetailLine name="Status">
                <StatusTypography>{pod.status}</StatusTypography>
            </DetailLine>
            <DetailLine name="Node">
                <LinkTypography>{pod.node}</LinkTypography>
            </DetailLine>
            <DetailLine name="Pod IP">
                <DetailTypography>{pod.podIP}</DetailTypography>
            </DetailLine>
            <DetailLine name="Service Account">
                <LinkTypography>{pod.serviceAccount}</LinkTypography>
            </DetailLine>
            <DetailLine name="QoS Class">
                <DetailTypography>{pod.qosClass}</DetailTypography>
            </DetailLine>
        </Box>
    )
}

function ContainerInfo(props) {
    const {container} = props;
    return (
        <Box
        sx={{ width:WIDTH,padding:"0px 10px"}}
    >
            <DetailLine name="Name">
                <DetailTypography>{container.name}</DetailTypography>
            </DetailLine>
            <DetailLine name="Status">
                <StatusTypography>{container.status}</StatusTypography>
            </DetailLine>
            <DetailLine name="Image">
                <DetailTypography>{container.image}</DetailTypography>
            </DetailLine>
            <DetailLine name="Ports">
                <ArrayTypography items={container.ports}></ArrayTypography>
            </DetailLine>
            <DetailLine name="Command">
                {
                    container.commands &&
                    container.commands.map((item,index) => {
                        return (
                            <DetailTypography key={index}>{item}</DetailTypography>
                        )
                    })

                }
            </DetailLine>
            <DetailLine name="Arguments">
                {
                    container.args &&
                    container.args.map((item,index) => {
                        return (
                            <DetailTypography key={index}>{item}</DetailTypography>
                        )
                    })
                }
            </DetailLine>
            </Box>
    )
}

function VolumeInfo() {
    return (
        <Box
        sx={{ width:WIDTH,padding:"0px 10px"}}
    >
            <h1>todo</h1>
        </Box>
    )
}


function TolerationsInfo() {
    return (
        <Box
        sx={{ width:WIDTH,padding:"0px 10px"}}
    >
           <h1>todo</h1>
            </Box>
    )
}


function createHandler(row) {
    return [
        CommonHandler.createDelete("pod", row.namespace, row.name),
        CommonHandler.createEdit("pod", row.namespace, row.name),
        CommonHandler.createLog("pod", row.namespace, row.name),
        CommonHandler.createShell(),
    ]
}

export default function PodDetail(props) {

    const defaultData = {
        name:"",
        namespace:"",
        created:"",
        labels:[],
        controlledBy:{
            kind:"",
            name:""
        },
        status:"",
        node:"",
        podIP:"",
        serviceAccount:"",
        qosClass:"",
        containers:[]
    };

    const {namespace,name,open,onClose} = props; 
    const [data, setData] = React.useState(defaultData);



    React.useEffect(() => {
        async function fetchData() {
            const result = await http.podDetail(namespace,name)
            setData(result);
            console.log(name)
        }
        if (open) {
            fetchData();
        }
    }, [open]);

    const info = () => (
       <>
            <DetailBar
                title={"Pod:" + data.name} 
                moreHandles={createHandler(data)}/>
            <Toolbar />
            <PodInfo pod={data}/>
            <TitleBar title="Containers" />
            {
                data.containers && data.containers.map((item,index) => {
                    return (
                        <ContainerInfo key={index} container={item}/>
                    )
                })
            }
            <TitleBar title="Volumes" />
            <VolumeInfo />
            <TitleBar title="Tolerations" />
            <TolerationsInfo />
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