import styled from "@emotion/styled";
import {Card, Col, Descriptions, DescriptionsProps, Row} from "antd";
import {useEffect, useRef, useState} from "react";
import {useRequest} from "ahooks";
import {cpuInfo, serverInfo} from "@/api/monitor.ts";
import {formatDuring} from "@/utils/time.ts";
import {Area, AreaConfig} from "@ant-design/charts";
import Monitor from "@/types/monitor.ts";

const MonitorServerPage = () => {


    const [data, setData] = useState<Monitor.CpuInfoResponse[]>([]);
    const [items, setItems] = useState<DescriptionsProps['items']>([])

    const serverRequest = useRequest(serverInfo, {
        manual: true,
        onSuccess: (data)=> {
            setItems([
                {
                    key: '1',
                    label: '操作系统',
                    children: data.os,
                },
                {
                    key: '2',
                    label: '系统版本',
                    children: data.platform,
                },
                {
                    key: '3',
                    label: '版本编码',
                    children: data.platformVersion,
                },
                {
                    key: '4',
                    label: 'CPU架构',
                    children: data.kernelArch,
                },
                {
                    key: '5',
                    label: '主机名称',
                    children: data.hostname,
                },
                {
                    key: '6',
                    label: '运行时长',
                    children: formatDuring(data.runTime),
                },
            ])
        }
    })

    const cpuRequest = useRequest(cpuInfo, {
        manual: true,
        onSuccess: (result)=> {
            setData([...data, {time: result.time, num: Number(result.num.toFixed(2))}]);
        }
    })

    let intervalRef = useRef<NodeJS.Timeout | null>(null);


    const config: AreaConfig = {
        data,
        xField: 'time',
        yField: 'num',
        yAxis: {
            range: [0, 100]
        },
        onReady: chart => {
            intervalRef.current = setInterval(async () => {
                let response = await cpuInfo();
                const newItem: Monitor.CpuInfoResponse = {time: response.time, num: Number(response.num.toFixed(2))};
                setData((origin) => {
                    return [...(origin.length > 10 ? origin.slice(1) : origin), newItem];
                })
                console.log("newData", response, data);
                chart.changeData(data)
                chart.render();
            }, 3000)
        }
    };



    useEffect(() => {
        serverRequest.run();
        cpuRequest.run();
        return () => {
            if (intervalRef.current !== null) {
                clearInterval(intervalRef.current);
            }
        }
    }, [])

    return <Container>
        <Card className="server-card server-info" title="基本信息" loading={serverRequest.loading}>
            <Descriptions items={items} />
        </Card>
        <Row gutter={5}>
            <Col span={12}>
                <Card className="server-card server-status" title="CPU">
                    <Area {...config} />
                </Card>
            </Col>
            <Col span={12}>
                <Card className="server-card server-status" title="内存">

                </Card>
            </Col>
            <Col span={12}>
                <Card className="server-card server-status" title="磁盘">

                </Card>
            </Col>
            <Col span={12}>
                <Card className="server-card server-status" title="流量统计">

                </Card>
            </Col>
        </Row>
    </Container>
}

const Container = styled.div`
    //background-color: #ffffff;
    //padding: 16px;
    //border-radius: 5px;
    .server-card {
        border-radius: 5px;
        margin-bottom: 5px;
        .ant-card-head {
            min-height: 40px;
        }
    }
    .server-status {
        
    }
    
    
`

// https://antv-g2.gitee.io/zh/examples/area/basic#basic
export default MonitorServerPage;
