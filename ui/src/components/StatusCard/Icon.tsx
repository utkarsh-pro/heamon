import { Status, COLORGREEN, COLORRED, COLORGRAY } from "../../global/constants"
import CheckCircleIcon from '@material-ui/icons/CheckCircle';
import CancelIcon from '@material-ui/icons/Cancel';
import HelpOutlineIcon from '@material-ui/icons/HelpOutline';

interface IconProps {
    status: Status;
    style?: React.CSSProperties;
}

function Icon({ status, style = {} }: IconProps) {
    if (status === Status.OK) return <CheckCircleIcon style={{ color: COLORGREEN, ...style }}/>
    if (status === Status.FAIL) return <CancelIcon style={{ color: COLORRED, ...style }} />
    
    return <HelpOutlineIcon style={{ color: COLORGRAY, ...style }} />
}

export default Icon
