import { Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Button } from "@mui/material";

interface ConfirmationModalProps {
    open: boolean;
    onClose: () => void;
    onConfirm: () => void;
    title: string;
    message: string;
    confirmText?: string;
    cancelText?: string;
    confirmColor?: "primary" | "secondary" | "error" | "info" | "success" | "warning";
    isLoading?: boolean;
}

export default function ConfirmationModal({
    open,
    onClose,
    onConfirm,
    title,
    message,
    confirmText = "Confirm",
    cancelText = "Cancel",
    confirmColor = "primary",
    isLoading = false
}: ConfirmationModalProps) {
    return (
        <Dialog
            open={open}
            onClose={onClose}
            aria-labelledby="confirmation-dialog-title"
            aria-describedby="confirmation-dialog-description"
        >
            <DialogTitle id="confirmation-dialog-title">
                {title}
            </DialogTitle>
            <DialogContent>
                <DialogContentText id="confirmation-dialog-description">
                    {message}
                </DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button onClick={onClose} color="primary" disabled={isLoading}>
                    {cancelText}
                </Button>
                <Button 
                    onClick={onConfirm} 
                    color={confirmColor} 
                    autoFocus
                    disabled={isLoading}
                >
                    {isLoading ? 'Processing...' : confirmText}
                </Button>
            </DialogActions>
        </Dialog>
    );
}
