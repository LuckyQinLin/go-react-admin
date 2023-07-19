import styled from "@emotion/styled";
import {Button} from "antd";
import Bg from "@/assets/images/403.svg";

const NotAuthPage = () => {
    return <Container>
        <div className="text-center">
            <img src={Bg} alt=""/>
        </div>
        <div className="text-center">
            <h1 className="text-base text-gray-500">抱歉，你无权访问该页面</h1>
            <Button type="primary">回到首页</Button>
        </div>
    </Container>
}

const Container = styled.div`
    display: flex;
    flex-direction: column;
    flex-wrap: nowrap;
    align-items: center;
    justify-content: center;
    width: 100%;
    border-radius: 4px;
    padding: 50px 0;
    
    .text-center {
        display: flex;
        flex-direction: column;
        flex-wrap: nowrap;
        align-items: center;
        h1 {
            color: #666;
            padding: 20px 0;
        }
    }
    
    img {
      width: 350px;
      margin: 0 auto;
    }
`

export default NotAuthPage;