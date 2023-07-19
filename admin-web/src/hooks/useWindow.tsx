import {useEffect, useState} from "react";

interface WindowSize {
	width: number;
	height: number;
}

const useWindowSize = (): WindowSize => {
	const [windowSize, setWindowSize] = useState<WindowSize>({
		width: 0,
		height: 0
	});

	const handler = () => {
		setWindowSize({
			width: window.innerWidth,
			height: window.innerHeight,
		})
	}

	useEffect(() => {

		handler();

		window.addEventListener('resize', handler);

		return () => {
			window.removeEventListener('resize', handler);
		}
	}, [])

	return windowSize
}

export default useWindowSize;