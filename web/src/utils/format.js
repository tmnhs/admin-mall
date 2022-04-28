import { formatTimeToStr } from '@/utils/date'
import { getDict } from '@/utils/dictionary'
import moment from 'moment'
export const formatBoolean = (bool) => {
  if (bool !== null) {
    return bool ? '是' : '否'
  } else {
    return ''
  }
}
export const formatDate = (time) => {
  if (time !== null && time !== '') {
    var date = new Date(time)
    return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
  } else {
    return ''
  }
}
//value 格式为13位unix时间戳
//10位unix时间戳可通过value*1000转换为13位格式
export const  formatUnixDate =(value) =>{
  return moment(value*1000).format('YYYY-MM-DD hh:mm:ss')
};


export const filterDict = (value, options) => {
  const rowLabel = options && options.filter(item => item.value === value)
  return rowLabel && rowLabel[0] && rowLabel[0].label
}

export const getDictFunc = async(type) => {
  const dicts = await getDict(type)
  return dicts
}
