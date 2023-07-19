import React from "react";
import {Button, ButtonProps} from "antd";
import styled from "@emotion/styled";


interface CustomBtnProp extends ButtonProps{
    custom: "Dust" | "Volcano" | "Sunset" | "Cyan" | "Green" | "Daybreak" | "Purple";
    title?: string;
}

export const CustomButton: React.FC<CustomBtnProp> = (props) => {

    switch (props.custom) {
        case "Dust":
            return <DustButton {...props} >{props.title}</DustButton>
        case "Volcano":
            return <VolcanoButton {...props} >{props.title}</VolcanoButton>
        case "Sunset":
            return <SunsetButton {...props} >{props.title}</SunsetButton>
        case "Cyan":
            return <CyanButton {...props} >{props.title}</CyanButton>
        case "Green":
            return <GreenButton {...props} >{props.title}</GreenButton>
        case "Daybreak":
            return <DaybreakButton {...props} >{props.title}</DaybreakButton>
        case "Purple":
            return <PurpleButton {...props} >{props.title}</PurpleButton>
    }

}

const DustButton = styled(Button)`
  background-color: #F5222D;
  border-color: #F5222D;

  &:hover, &:focus {
    background-color: #ff4d4f;
    border-color: #ff4d4f;
  }

  &:active, &.active {
    background-color: #cf1322;
    border-color: #cf1322;
  }
`
const VolcanoButton = styled(Button)`
  background-color: #FA541C;
  border-color: #FA541C;

  &:hover, &:focus {
    background-color: #ff7a45;
    border-color: #ff7a45;
  }

  &:active, &.active {
    background-color: #d4380d;
    border-color: #d4380d;
  }
`

const SunsetButton = styled(Button)`
  background-color: #FAAD14;
  border-color: #FAAD14;

  &:hover, &:focus {
    background-color: #ffc53d;
    border-color: #ffc53d;
  }

  &:active, &.active {
    background-color: #d48806;
    border-color: #d48806;
  }
`

const CyanButton = styled(Button)`
  background-color: #13C2C2;
  border-color: #13C2C2;

  &:hover, &:focus {
    background-color: #36cfc9;
    border-color: #36cfc9;
  }

  &:active, &.active {
    background-color: #08979c;
    border-color: #08979c;
  }  
`

const GreenButton = styled(Button)`
  background-color: #52C41A;
  border-color: #52C41A;

  &:hover, &:focus {
    background-color: #73d13d;
    border-color: #73d13d;
  }

  &:active, &.active {
    background-color: #389e0d;
    border-color: #389e0d;
  }
`

const DaybreakButton = styled(Button)`
  background-color: #1890FF;
  border-color: #1890FF;

  &:hover, &:focus {
    background-color: #096dd9;
    border-color: #096dd9;
  }

  &:active, &.active {
    background-color: #40a9ff;
    border-color: #40a9ff;
  }
`

const PurpleButton = styled(Button)`
  background-color: #722ED1;
  border-color: #722ED1;

  &:hover, &:focus {
    background-color: #9254de;
    border-color: #9254de;
  }

  &:active, &.active {
    background-color: #531dab;
    border-color: #531dab;
  }
`


