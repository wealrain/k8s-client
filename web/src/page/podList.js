import * as React from 'react';
import PageinationTable from '../component/PageinationTable';
import CommonHandler from './common'
import { timeDifference } from '../util';
import list from '../http/list';
import DataFilter from '../http/dataFilter';
import PodDetail from '../component/detail/PodDetail';
import { AppContext } from '../App';
import SearchBar from '../component/SearchBar';


const columns = ["status","name","namespace", "node", "restarts", "age"]   

function createHandler(row) {
    return [
        CommonHandler.createDelete("pod", row.namespace, row.name),
        CommonHandler.createEdit("pod", row.namespace, row.name),
        CommonHandler.createLog("pod", row.namespace, row.name),
        CommonHandler.createShell(),
    ]
}


function PodList() {
    const {cluster,namespace} = React.useContext(AppContext);
    const [data, setData] = React.useState([]);
    const [total, setTotal] = React.useState(0);
    const [current, setCurrent] = React.useState(0); // MUI page start from 0
    const [loading, setLoading] = React.useState(false);
    const [openDetail, setOpenDetail] = React.useState(false);
    const [podDetail, setPodDetail] = React.useState({});
    const [searchName, setSearchName] = React.useState("");
    // 创建数据过滤器
    const dataFilter = new DataFilter();

    async function fetchData() {
        setLoading(true);
        dataFilter.setNameFilter(searchName);
        dataFilter.setPage(current + 1);
        const result = await list.listPods(namespace,dataFilter.toJson());
        if(result.total === 0) {
            setData([]);
            setTotal(0);
            setLoading(false);
            return;
        }
        setData(result.list.map(item => {
            return {
                status: item.status,
                name: item.name,
                namespace: item.namespace,
                node: item.node,
                restarts: item.restarts,
                age: timeDifference(item.age)
                }
        }));
        setTotal(result.total);
        setCurrent(result.pageNum - 1);
        setLoading(false);
    }
    
    React.useEffect(() => {
        fetchData("");
    }, [current,cluster,namespace,searchName]);

    return (
        <>
            <SearchBar onSearchResource={(value)=>{
                setCurrent(0);
                setSearchName(value);
            }}/>
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
                    setPodDetail(row);
                }}
            />
            <PodDetail
                namespace= {podDetail.namespace}
                name={podDetail.name}
                open={openDetail} 
                onClose={() => setOpenDetail(false)} 
            />
        </>
    );
}

export default PodList;