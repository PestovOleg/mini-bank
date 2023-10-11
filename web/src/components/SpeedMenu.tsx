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
import { Fab } from '@mui/material';
import { red } from '@mui/material/colors';

import PaymentIcon from '@mui/icons-material/Payment';
import AddCardIcon from '@mui/icons-material/AddCard';
import NewAccountDialog from './NewAccount';
import PaymentDialog from './Payment';




export default function SpeedDialMenu() {
    const [open, setOpen] = React.useState(false);
    const [newAccountOpen, setNewAccountOpen] = React.useState(false);
    const [paymentOpen, setPaymentOpen] = React.useState(false);

    const handleOpen = () => setOpen(true);
    const handleClose = () => setOpen(false);
    const handleNewAccountDialogOpen = () => { 
        handleClose(); 
        setNewAccountOpen(true);        
    };

    const handleNewAccountDialogClose = () => { // закрыть диалоговое окно
        setNewAccountOpen(false);
    };

    const handlePaymentDialogOpen = () => { 
        handleClose(); 
        setPaymentOpen(true);        
    };

    const handlePaymentDialogClose = () => { // закрыть диалоговое окно
        setPaymentOpen(false);
    };

    const actions = [
        { icon: <AddCardIcon />, name: 'Открыть счет',action: handleNewAccountDialogOpen},
        { icon: <PaymentIcon />, name: 'Переводы', action: handlePaymentDialogOpen },
    ];
    return (
        <Box sx={{
            height: 220,
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
                        tooltipOpen
                        onClick={() => {
                            action.action();
                          }}
                    />
                ))}
            </SpeedDial>
            <NewAccountDialog open={newAccountOpen} setOpen={setNewAccountOpen} />
            <PaymentDialog open={paymentOpen} setOpen={setPaymentOpen} />

        </Box>
    );
}