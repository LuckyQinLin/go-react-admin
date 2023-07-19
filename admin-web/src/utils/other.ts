/**
 * @description: 判断值是否未某个类型
 */
export function is(val: unknown, type: string) {
	return toString.call(val) === `[object ${type}]`;
}

export function isFunction<T = any>(val: unknown): val is T {
	return is(val, 'Function') || is(val, 'AsyncFunction');
}

/**
 * 文件大小转换
 * @param fileSizeInBytes
 */
export function formatFileSize(fileSizeInBytes: number): string {
	if (fileSizeInBytes < 1024) {
		return fileSizeInBytes + " B";
	} else if (fileSizeInBytes < 1024 * 1024) {
		const fileSizeInKB = fileSizeInBytes / 1024;
		return fileSizeInKB.toFixed(2) + " KB";
	} else if (fileSizeInBytes < 1024 * 1024 * 1024) {
		const fileSizeInMB = fileSizeInBytes / (1024 * 1024);
		return fileSizeInMB.toFixed(2) + " MB";
	} else {
		const fileSizeInGB = fileSizeInBytes / (1024 * 1024 * 1024);
		return fileSizeInGB.toFixed(2) + " GB";
	}
}