import * as React from 'react';
import PageinationTable from '../component/PageinationTable';
import CommonHandler from './common'
import { timeDifference } from '../util';
import list from '../http/list';
import DataFilter from '../http/dataFilter';
import DeploymentDetail from '../component/detail/DeploymentDetail';

const columns = ["name","pods", "replicas", "age"]   

function createHandler(row) {
    return [
        CommonHandler.createDelete("deployment", row.namespace, row.name),
        CommonHandler.createEdit("deployment", row.namespace, row.name),
        CommonHandler.createScale(),
        CommonHandler.createRestart(),
        CommonHandler.createPause(row.namespace,row.name),
        CommonHandler.createResume(row.namespace,row.name),      
    ]
}

function DeploymentList() {
    const [data, setData] = React.useState([]);
    const [total, setTotal] = React.useState(0);
    const [current, setCurrent] = React.useState(0); // MUI page start from 0
    const [loading, setLoading] = React.useState(false);
    const [openDetail, setOpenDetail] = React.useState(false);
    const [detail, setDetail] = React.useState({});

    // 创建数据过滤器
    const dataFilter = new DataFilter();

    async function fetchData() {
        setLoading(true);
        dataFilter.setPage(current + 1);
        const result = await list.listDeployments('wuxi-dev',dataFilter.toJson());
        setData(result.list.map(item => {
            return {
                name: item.name,
                namespace: item.namespace,
                pods: item.pods,
                replicas: item.replicas,
                age: timeDifference(item.age)
                }
        }));
        setTotal(result.total);
        setCurrent(result.pageNum - 1);
        setLoading(false);
    }
    
    React.useEffect(() => {
        fetchData();
    }, [current]);

    return (
        <>
        <PageinationTable 
            columns={columns} 
            data={data} 
            loading={loading}
            moreHandler={createHandler}
            total={total}
            current={current}
            onChange={(e,page) => {
                setCurrent(page);
            }}
            handleClick={(e,row) => {
                setOpenDetail(true);
                setDetail(row);
            }}
        />
        <DeploymentDetail
                namespace= {detail.namespace}
                name={detail.name}
                open={openDetail} 
                onClose={() => setOpenDetail(false)} 
            />
        </>

    );
}

export default DeploymentList;