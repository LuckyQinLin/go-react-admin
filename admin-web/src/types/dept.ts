namespace Dept {
    // 用户部门信息
    export interface UserDeptProp {
        deptId: number;  // 岗位ID
        parentId: number; // 上次岗位
        deptName: string; // 岗位名称
        deptPath: string; // 岗位路径
    }
}

export default Dept;
