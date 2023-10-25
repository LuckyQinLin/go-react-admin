import styled from "@emotion/styled";
import {Card, Col, Descriptions, DescriptionsProps, Row} from "antd";
import {useEffect, useState} from "react";
import {useRequest} from "ahooks";
import {serverInfo} from "@/api/monitor.ts";
import {formatDuring} from "@/utils/time.ts";
import CpuUsageChart from "@/pages/monitor/server/CpuUsageChart.tsx";
import MemUsageChart from "@/pages/monitor/server/MemUsageChart.tsx";

const MonitorServerPage = () => {

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



    useEffect(() => {
        serverRequest.run();
    }, [])

    return <Container>
        <Card className="server-card server-info" title="基本信息" loading={serverRequest.loading}>
            <Descriptions items={items} />
        </Card>
        <Row gutter={5}>
            <Col span={12}>
                <CpuUsageChart />
            </Col>
            <Col span={12}>
                <MemUsageChart />
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
