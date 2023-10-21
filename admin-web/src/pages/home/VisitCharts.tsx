import {Column} from "@ant-design/charts";

const VisitCharts = () => {
	const data = [
		{
			type: '一月',
			sales: 38,
		},
		{
			type: '二月',
			sales: 52,
		},
		{
			type: '三月',
			sales: 61,
		},
		{
			type: '四月',
			sales: 105,
		},
		{
			type: '五月',
			sales: 48,
		},
		{
			type: '六月',
			sales: 38,
		},
		{
			type: '七月',
			sales: 38,
		},
		{
			type: '八月',
			sales: 67,
		},
		{
			type: '九月',
			sales: 38,
		},
		{
			type: '十月',
			sales: 90,
		},
		{
			type: '十一月',
			sales: 66,
		},
		{
			type: '十二月',
			sales: 38,
		},
	];

	const config = {
		data,
		xField: 'type',
		yField: 'sales',
		xAxis: {
			label: {
				autoHide: true,
				autoRotate: false,
			},
		},
		meta: {
			type: {
				alias: '月份',
			},
			sales: {
				alias: '访问数',
			},
		},
		minColumnWidth: 20,
		maxColumnWidth: 20,
	};

	return <Column {...config} />;
}

export default VisitCharts;