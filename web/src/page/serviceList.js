import * as React from 'react';
import PageinationTable from '../component/PageinationTable';
import CommonHandler from './common'
import { timeDifference } from '../util';
import list from '../http/list';
import DataFilter from '../http/dataFilter';

const columns = ["name","type", "clusterIP","externalIP","ports", "age"]   

function createHandler(row) {
    return [
        CommonHandler.createDelete("endpoint", row.name, row.namespace),
        CommonHandler.createEdit(),      
    ]
}

function ServiceList() {
    const [data, setData] = React.useState([]);
    const [total, setTotal] = React.useState(0);
    const [current, setCurrent] = React.useState(0); // MUI page start from 0
    const [loading, setLoading] = React.useState(false);

    // 创建数据过滤器
    const dataFilter = new DataFilter();

    async function fetchData() {
        setLoading(true);
        dataFilter.setPage(current + 1);
        const result = await list.listServices('wuxi-dev',dataFilter.toJson());
        console.log(result)
        setData(result.list.map(item => {
            return {
                name: item.name,
                type: item.type,
                clusterIP: item.clusterIP,
                externalIP: item.externalIP,
                ports: item.ports,
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

export default ServiceList;