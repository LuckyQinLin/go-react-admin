import {Area, AreaConfig} from "@ant-design/charts";
import {Card} from "antd";
import styled from "@emotion/styled";
import {cpuInfo} from "@/api/monitor.ts";
import Monitor from "@/types/monitor.ts";
import {useEffect, useRef} from "react";

const datas: Monitor.CpuInfoResponse[] = [];

const CpuUsageChart = () => {

    let intervalRef = useRef<NodeJS.Timeout | null>(null);

    const config: AreaConfig = {
        data: datas,
        xField: 'time',
        yField: 'num',
        xAxis: {
            range: [0, 1],
        },
        onReady: chart => {
            intervalRef.current = setInterval(async () => {
                let response = await cpuInfo();
                const newItem: Monitor.CpuInfoResponse = {time: response.time, num: Number(response.num.toFixed(2))};
                if (datas.length > 10) {
                    datas.shift();
                }
                datas.push(newItem)
                chart.changeData(datas)
            }, 3000)
        }
    };

    const getCpuData = async () => {
        let response = await cpuInfo();
        datas.push(response);
        // setData([response])
    }

    useEffect(() => {
        getCpuData()
        return () => {
            datas.splice(0, datas.length)
            if (intervalRef.current) {
                clearInterval(intervalRef.current)
            }
        }
    }, []);


    return <Container title="CPU">
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

export default CpuUsageChart;
