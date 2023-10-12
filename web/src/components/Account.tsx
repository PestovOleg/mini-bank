import React from 'react';
import { Box, Typography, List, ListItem, ListItemAvatar, Avatar, ListItemText, IconButton, Paper, Divider, ListItemIcon, Collapse, ListItemButton, Tooltip } from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';
import { IAccount } from '../models/types';
import AttachMoneyIcon from '@mui/icons-material/AttachMoneySharp';
import CurrencyRubleOutlined from '@mui/icons-material/CurrencyRubleOutlined';
import { ExpandLess, ExpandMore } from '@mui/icons-material';
import ModeIcon from '@mui/icons-material/ModeSharp';
import DeleteForeverSharpIcon from '@mui/icons-material/DeleteForeverSharp';
import ChangeAccountNameDialog from './ChangeAccountName';
import CloseAccountDialog from './CloseAccount';

// Определение пропсов
interface AccountProps {
    title: string;
    accounts: IAccount[];
}

const Account: React.FC<AccountProps> = ({ title, accounts }) => {
    const [open, setOpen] = React.useState(
        Array(accounts.length).fill(false)
    );
    const [changeNameDialogOpen, setChangeNameDialogOpen] = React.useState(false);
    const [closeAccountDialogOpen, setCloseAccountDialogOpen] = React.useState(false);
    const [currentAccount, setCurrentAccount] = React.useState<IAccount | null>(null);


    const openChangeNameDialog = (account: IAccount) => {
        setCurrentAccount(account);
        setChangeNameDialogOpen(true);
    };
    
    const openCloseAccountDialog = (account: IAccount) => {
        setCurrentAccount(account);
        setCloseAccountDialogOpen(true);
    };

    const closeCloseAccountDialog = (account: IAccount) => {        
        setCloseAccountDialogOpen(false);
    };

    const closeChangeNameDialog = () => {
        setChangeNameDialogOpen(false);
    };

    const handleClick = (index: number) => {
        const newOpen = [...open];
        newOpen[index] = !newOpen[index];
        setOpen(newOpen);
    };

    const sortedAccounts = React.useMemo(() => {
        return [...accounts].sort((a, b) => {
            const currencyCompare = a.currency.localeCompare(b.currency);
            if (currencyCompare !== 0) {
                return currencyCompare;
            }
            return b.amount - a.amount;
        });
    }, [accounts]);


    return (
        <Box sx={{ width: '100%',marginBottom:8 }}>
            <Typography variant="h6" component="div" sx={{ textAlign: 'center' }}>
                {title}
            </Typography>
            <div >
                <List>
                    {sortedAccounts.map((item, index) => (
                        <div key={index}>
                            <ListItem sx={{ boxShadow: '0px 3px 5px rgba(0, 0, 0, 0.2)', mb: 1 }}
                                secondaryAction={
                                    <ListItemButton onClick={() => handleClick(index)}>
                                        {open[index] ? <ExpandLess /> : <ExpandMore />}
                                    </ListItemButton>
                                }
                            >
                                <ListItemAvatar>
                                    <Avatar >
                                        {
                                            item.currency === '810' ? (<img alt="Ruble" src="/ruble.png" />) :
                                                item.currency === '840' ? (<img alt="Dollar" src="/dollar.png" />) : null
                                        }
                                    </Avatar>
                                </ListItemAvatar>
                                <ListItemText
                                    primary={item.name}
                                    secondary={item.amount.toLocaleString('ru-Ru', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}
                                />

                            </ListItem>
                            <Collapse in={open[index]} timeout="auto" unmountOnExit>
                                <List component="div" disablePadding>
                                    <ListItemButton sx={{ display: 'flex', justifyContent: 'space-around' }}>
                                        <ListItemText sx={{ ml: 7 }} primary={item.account} secondary={item.interestRate * 100 + '%'} />
                                        <Box sx={{ display: 'flex', flexDirection: 'column', alignItems: 'center', mr: 1 }}>
                                            <Tooltip title="Сменить имя счета">
                                                <IconButton size='small' onClick={() => { openChangeNameDialog(item) }}>
                                                    <ModeIcon />
                                                </IconButton>
                                            </Tooltip>
                                            <Tooltip title="Закрыть счет">
                                            <IconButton size='small' onClick={() => { openCloseAccountDialog(item) }}>
                                                    <DeleteIcon />
                                                </IconButton>
                                            </Tooltip>
                                        </Box>
                                    </ListItemButton>
                                </List>
                            </Collapse>
                        </div>


                    ))}
                </List>
            </div>
            <ChangeAccountNameDialog
                open={changeNameDialogOpen}
                setOpen={setChangeNameDialogOpen}
                account={currentAccount}
            />
            <CloseAccountDialog
                open={closeAccountDialogOpen}
                setOpen={setCloseAccountDialogOpen}
                account={currentAccount}
            />
        </Box>
    );
};

export default Account;
