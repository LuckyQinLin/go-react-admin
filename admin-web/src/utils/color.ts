const colors: string[] = ["red", "volcano", "gold", "magenta", "lime", "green", "cyan", "blue", "geekblue", "purple"];

export const color = (index: number) => {
    return colors[index % colors.length];
}