// 转化时间单位
const TimeUnit: string[] = ["年", "月", "天", "小时", "分钟", "刚刚"]
// 转化时间单位对应的时间长度（秒单位）
const TimeLen: number[] = [31536000, 2592000, 604800, 86400, 3600, 60, 1]

/**
 * 时间单位转换映射
 */
const TimeMapping: Map<string, number> = new Map([
	["年", 31536000],
	["月", 2592000],
	["星期", 604800],
	["天", 86400],
	["小时", 3600],
	["分钟", 60],
	["刚刚", 1],
])

/**
 * 时间转化
 * @param time
 * @constructor
 */
export const TimeConvert = (time: string) => {
	// 时间差值
	const timeValue: number = Math.round(new Date().getTime() / 1000) - new Date(time).getTime() / 1000;
	for (const [key, value] of TimeMapping) {
		const tmp = Math.floor(timeValue / value)
		if (tmp != 0) {
			return tmp + key + "前"
		}
	}

}

/**
 * 节流
 * @param {*} fn 将执行的函数
 * @param {*} time 节流规定的时间
 */
export const throttle = (fn: Function, time: number) => {
	let timer: any = null
	return (...args: any) => {
		// 若timer === false，则执行，并在指定时间后将timer重制
		if(!timer){
			fn.apply(this, args)
			timer = setTimeout(() => {
				timer = null
			}, time)
		}
	}
}

/**
 * 将时间戳转为 => 几天几小时几分钟几秒
 * @param mss
 */
export const formatDuring = (mss: number): string => {
	let days = Math.trunc(mss / (1000 * 60 * 60 * 24));
	let hours = Math.trunc((mss % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
	let minutes = Math.trunc((mss % (1000 * 60 * 60)) / (1000 * 60));
	let seconds = ((mss % (1000 * 60)) / 1000).toFixed(0);
	return days + " 天 " + hours + " 小时 " + minutes + " 分钟 " + seconds + " 秒 ";
}
