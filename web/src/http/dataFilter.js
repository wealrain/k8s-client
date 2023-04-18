class DataFilter {
    constructor() {
        this._resource = null;
        this._filters = [];
        this._sort = null;
        this._page = null;
    }

    setResource(resource) {
        this._resource = resource;
    }

    addFilter(propertyName,filterType,value) {
        this._filters.push({
            propertyName: propertyName,
            filterType: filterType,
            value: value
        });
    }

    setSort(propertyName,ascending) {
        this._sort = {
            propertyName: propertyName,
            ascending: ascending
        };
    }

    setPage(pageNum,pageSize = 10) {
        this._page = {
            pageNum: pageNum,
            pageSize: pageSize
        };
    }

    toJson() {
        let json = {};
        if (this._resource) {
            json.resource = this._resource;
        }
        if (this._filters.length > 0) {
            json.filters = this._filters;
        }
        if (this._sort) {
            json.sort = this._sort;
        }
        if (this._page) {
            json.page = this._page;
        }
        return json;
    }

    
}

// 定义过滤器类型
DataFilter.FilterType = {
    Equal: 0,
    Contains: 1,
    GreaterThan: 2,
    GreaterThanOrEqual: 3,
    LessThan: 4,
    LessThanOrEqual: 5,
    Between: 6,
};
    


export default DataFilter;