import * as React from 'react';
import Box from '@mui/material/Box';
import Backdrop from '@mui/material/Backdrop';
import SpeedDial from '@mui/material/SpeedDial';
import SpeedDialIcon from '@mui/material/SpeedDialIcon';
import SpeedDialAction from '@mui/material/SpeedDialAction';
import PaymentIcon from '@mui/icons-material/Payment';
import AddCardIcon from '@mui/icons-material/AddCard';
import NewAccountDialog from './NewAccount';
import PaymentDialog from './Payment';
import TopupDialog from './TopUp';
import '../styles.css';
import RedeemIcon from '@mui/icons-material/Redeem';
import store from '../store/store';




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

    const handlePaymentDialogOpen = () => {
        handleClose();
        setPaymentOpen(true);
    };

    const handleTopupDialogOpen = () => {
        handleClose();
        setTopupOpen(true);
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