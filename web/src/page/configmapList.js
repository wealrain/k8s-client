import * as React from 'react';
import PageinationTable from '../component/PageinationTable';
import CommonHandler from './common';
import { timeDifference } from '../util';
import list from '../http/list';
import DataFilter from '../http/dataFilter';
import ConfigmapDetail from '../component/detail/ConfigmapDetail';
import { AppContext } from '../App';
import SearchBar from '../component/SearchBar';

const columns = ["name","namespace","keys", "age"]   

function createHandler(row) {
    return [
        CommonHandler.createDelete("configmap", row.name, row.namespace),
        CommonHandler.createEdit(),         
    ]
}

function ConfigMapList() {
    const {cluster,namespace} = React.useContext(AppContext);
    const [data, setData] = React.useState([]);
    const [total, setTotal] = React.useState(0);
    const [current, setCurrent] = React.useState(0); // MUI page start from 0
    const [loading, setLoading] = React.useState(false);
    const [openDetail, setOpenDetail] = React.useState(false);
    const [detail, setDetail] = React.useState({});
    const [searchName, setSearchName] = React.useState("");
    // 创建数据过滤器
    const dataFilter = new DataFilter();

    async function fetchData() {
        setLoading(true);
        dataFilter.setNameFilter(searchName);
        dataFilter.setPage(current + 1);
        const result = await list.listConfigmaps(namespace,dataFilter.toJson());
        setData(result.list.map(item => {
            return {
                name: item.name,
                keys: item.keys,
                namespace: item.namespace,
                age: timeDifference(item.age)
            }
        }));
        setTotal(result.total);
        setCurrent(result.pageNum - 1);
        setLoading(false);
    }
    
    React.useEffect(() => {
        fetchData();
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
                setDetail(row);
            }}
        />
        <ConfigmapDetail
                namespace= {detail.namespace}
                name={detail.name}
                open={openDetail} 
                onClose={() => setOpenDetail(false)} 
            />
        </>
    );
}

export default ConfigMapList;