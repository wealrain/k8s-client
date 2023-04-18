// 将求当前时间和时间戳的差值 
export const timeDifference = (time) => {
    let now = new Date().getTime() / 1000;
    let diff = now - time;
    let result = null;
    if (diff < 0) {
        return;
    }
    let min = diff / 60;
    let hour = min / 60;
    let day = hour / 24;
    let month = day / 30;
    let year = month / 12;
    if (year >= 1) {
        result = parseInt(year) + "年前";
    } else if (month >= 1) {
        result = parseInt(month) + "个月前";
    } else if (day >= 1) {
        result = parseInt(day) + "天前";
    } else if (hour >= 1) {
        result = parseInt(hour) + "小时前";
    } else if (min >= 1) {
        result = parseInt(min) + "分钟前";
    } else {
        result = "刚刚";
    }
    return result;
}