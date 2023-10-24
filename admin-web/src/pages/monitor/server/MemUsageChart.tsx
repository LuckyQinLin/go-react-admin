import {Card} from "antd";
import styled from "@emotion/styled";
import {memInfo} from "@/api/monitor.ts";
import Monitor from "@/types/monitor.ts";
import {useEffect, useRef} from "react";
import {Area, AreaConfig} from "@ant-design/charts";

const datas: Monitor.MemAreaField[] = []

const MemUsageChart = () => {

    // const [data, setData] = useState<Monitor.MemUsageResponse[]>([]);
    let intervalRef = useRef<NodeJS.Timeout | null>(null);

    const config: AreaConfig = {
        data: datas,
        xField: 'time',
        yField: 'data',
        seriesField: 'types',
        xAxis: {
            range: [0, 1],
        },
        onReady: chart => {
            intervalRef.current = setInterval(async () => {
                let response = await memInfo();
                if (datas.length > 40) {
                    datas.splice(0, 4);
                }
                datas.push(...dataConvert(response))
                chart.changeData(datas)
            }, 3000)
        }
    };

    const dataConvert = (data: Monitor.MemUsageResponse): Monitor.MemAreaField[] => {
        return [
            {
                data: Number((data.total / 1024 / 1024 / 1024).toFixed(2)),
                types: '总内存',
                time: data.time,
            },
            {
                data: Number((data.used / 1024 / 1024 / 1024).toFixed(2)),
                types: '占用',
                time: data.time,
            },
            {
                data: Number((data.free / 1024 / 1024 / 1024).toFixed(2)),
                types: '空闲',
                time: data.time,
            },
            {
                data: Number(data.percent.toFixed(2)),
                types: '占比',
                time: data.time,
            },
        ]
    }

    const getMemData = async () => {
        let response = await memInfo();
        datas.push(...dataConvert(response))
    }

    useEffect(() => {
        getMemData();
        return () => {
            datas.splice(0, datas.length);
            if (intervalRef.current) {
                clearInterval(intervalRef.current)
            }
        }
    }, []);


    return <Container title="内存">
        <Area {...config} />
    </Container>
}

const Container = styled(Card)`
  border-radius: 5px;
  margin-bottom: 5px;
  .ant-card-head {
    min-height: 40px;
  }
`

export default MemUsageChart;
