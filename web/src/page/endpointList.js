import * as React from 'react';
import PageinationTable from '../component/PageinationTable';
import CommonHandler from './common';
import { timeDifference } from '../util';
import list from '../http/list';
import DataFilter from '../http/dataFilter';

const columns = ["name","endpoints", "age"]   

function createHandler(row) {
    return [
        CommonHandler.createDelete("endpoint", row.name, row.namespace),
        CommonHandler.createEdit(),      
    ]
}

function EndpointList() {
    const [data, setData] = React.useState([]);
    const [total, setTotal] = React.useState(0);
    const [current, setCurrent] = React.useState(0); // MUI page start from 0
    const [loading, setLoading] = React.useState(false);

    // 创建数据过滤器
    const dataFilter = new DataFilter();

    async function fetchData() {
        setLoading(true);
        dataFilter.setPage(current + 1);
        const result = await list.listEndpoints('wuxi-dev',dataFilter.toJson());
        setData(result.list.map(item => {
            return {
                name: item.name,
                endpoints: item.endpoints,
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
        />
    );
}

export default EndpointList;