import styled from "@emotion/styled";
import {Card, Col, Row} from "antd";
import {useEffect} from "react";
import {useRequest} from "ahooks";
import {serverInfo} from "@/api/monitor.ts";

const MonitorServerPage = () => {

    const {loading, run} = useRequest(serverInfo, {
        manual: true,
        onSuccess: (data)=> {
            console.log(data)
        }
    })


    useEffect(() => {
        run()
    }, [])

    return <Container>
        <Card className="server-card server-info" title="基本信息" loading={loading}>

        </Card>
        <Row gutter={5}>
            <Col span={12}>
                <Card className="server-card server-status" title="CPU">

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