import * as React from 'react';
import Box from '@mui/material/Box';
import Backdrop from '@mui/material/Backdrop';
import SpeedDial from '@mui/material/SpeedDial';
import SpeedDialIcon from '@mui/material/SpeedDialIcon';
import SpeedDialAction from '@mui/material/SpeedDialAction';
import FileCopyIcon from '@mui/icons-material/FileCopyOutlined';
import SaveIcon from '@mui/icons-material/Save';
import PrintIcon from '@mui/icons-material/Print';
import ShareIcon from '@mui/icons-material/Share';
import styled from '@emotion/styled';
import { makeStyles, useTheme } from '@mui/material/styles';
import { red } from '@mui/material/colors';
import PaymentIcon from '@mui/icons-material/Payment';
import AddCardIcon from '@mui/icons-material/AddCard';
import NewAccountDialog from './NewAccount';
import PaymentDialog from './Payment';
import TopupDialog from './TopUp';
import '../styles.css';
import RedeemIcon from '@mui/icons-material/Redeem';
import store from '../store/store';
import { Height } from '@mui/icons-material';




export default function SpeedDialMenu() {
    const [open, setOpen] = React.useState(false);
    const [newAccountOpen, setNewAccountOpen] = React.useState(false);
    const [paymentOpen, setPaymentOpen] = React.useState(false);
    const [topupOpen, setTopupOpen] = React.useState(false);

    const handleOpen = () => setOpen(true);
    const handleClose = () => setOpen(false);
    const handleNewAccountDialogOpen = () => {
        handleClose();
        setNewAccountOpen(true);
    };

    const handleNewAccountDialogClose = () => {
        setNewAccountOpen(false);
    };

    const handlePaymentDialogOpen = () => {
        handleClose();
        setPaymentOpen(true);
    };

    const handlePaymentDialogClose = () => {
        setPaymentOpen(false);
    };

    const handleTopupDialogOpen = () => {
        handleClose();
        setTopupOpen(true);
    };

    const handleTopupDialogClose = () => {
        setTopupOpen(false);
    };

    const createAccountToggle = store.toggleStore.getFeature("CreateAccountToggle")
    const topupToggle = store.toggleStore.getFeature("TopUpToggle")
    const WithdrawToggle = store.toggleStore.getFeature("WithdrawToggle")

    const createAccountToggleNumber = createAccountToggle ? 1 : 0
    const topupToggleNumber = topupToggle ? 1 : 0
    const WithdrawToggleNumber = WithdrawToggle ? 1 : 0

    const boxHeight = () => {
        const countToggles = createAccountToggleNumber + topupToggleNumber + WithdrawToggleNumber
        if (countToggles === 3) { return '275px' }
        else if (countToggles === 2) { return '220px' }
        else if (countToggles === 1) { return '160px' }
        else return '105px'
    }

    interface Action {
        icon: JSX.Element;
        name: string;
        action: () => void;
    }
    const actions: Action[] = [];

    if (createAccountToggle) {
        actions.push({ icon: <AddCardIcon />, name: 'Открыть счет', action: handleNewAccountDialogOpen },)
    } 

    if (topupToggle) {
        actions.push({ icon: <PaymentIcon />, name: 'Пополнение', action: handleTopupDialogOpen },)
    }

    if (WithdrawToggle) {
        actions.push({ icon: <RedeemIcon />, name: 'Переводы', action: handlePaymentDialogOpen },)
    }

    return (

        <Box sx={{
            height: boxHeight(),
            position: 'absolute',

            bottom: '0px', // position it at the bottom of the AppBar        
            zIndex: 1300, // make sure it appears above the AppBar
            flexGrow: 1,
            transform: 'translateZ(0px)',
        }}>
            <Backdrop open={open} invisible={true} />

            <SpeedDial
                ariaLabel="SpeedDial tooltip example"
                sx={{ position: 'relative' }}
                icon={<SpeedDialIcon />}
                FabProps={{
                    sx: {
                        bgcolor: 'secondary.main',
                        '&:hover': {
                            bgcolor: 'secondary.main',
                        }
                    }
                }}
                onClose={handleClose}
                onOpen={handleOpen}
                open={open}
            >
                {actions.map((action) => (

                    <SpeedDialAction
                        key={action.name}
                        icon={action.icon}
                        tooltipTitle={action.name}
                        //TooltipClasses={{tooltip: '.MuiSpeedDialAction-staticTooltipLabel'}}
                        FabProps={{
                            sx: {
                                bgcolor: 'main.main',
                                '&:hover': {
                                    bgcolor: 'gray.main',
                                }
                            }
                        }}
                        tooltipOpen
                        onClick={() => {
                            action.action();
                        }}
                    />

                ))}
            </SpeedDial>
            <NewAccountDialog open={newAccountOpen} setOpen={setNewAccountOpen} />
            <PaymentDialog open={paymentOpen} setOpen={setPaymentOpen} />
            <TopupDialog open={topupOpen} setOpen={setTopupOpen} />
        </Box>
    );
}