// import React, {useEffect, useState} from "react";
// import {Button} from "antd";
// import {ButtonProps} from "antd/es/button/button";
// import {permsKeys} from "@/pages/layout";
//
// // 扩展属性
// type WithTitle<T, P = string> = T & { title: P, code: P };
//
// const AuthButton: React.FC<WithTitle<ButtonProps>> = (props) => {
//
//     const [isShow, setIsShow] = useState<boolean>(true);
//     const userStore = useSelector((state) => state.user);
//
//     useEffect(() => {
//         // console.log("props", props);
//         // if (userStore.perms) {
//         //     let keys = permsKeys(userStore.perms);
//         //     const flag = keys.includes(props.code)
//         //     console.log("AuthButton -> true", keys, flag, props.code);
//         //     setIsShow(flag)
//         // } else {
//         //     console.log("AuthButton -> false")
//         //     setIsShow(false);
//         // }
//         setIsShow(userStore.perms ? permsKeys(userStore.perms).includes(props.code) : false)
//
//     }, [])
//
//
//
//     return isShow ? <Button {...props}>{props.title}</Button> : null;
// }
//
// export default AuthButton;
