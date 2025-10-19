'use client';

import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import {
    Box,
    Button,
    TextField,
    Switch,
    FormControlLabel,
    Typography,
    Paper,
    Alert,
} from '@mui/material';
import { useState } from 'react';
import { en } from 'zod/locales';

// Zod schema for route validation
const routeSchema = z.object({
    name: z.string().min(1, 'Name is required').min(3, 'Name must be at least 3 characters'),
    host: z.string()
        .min(1, 'Host is required')
        .regex(/^[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/, 'Invalid host format (e.g., example.com)'),
    backend: z.string().url('Backend must be a valid URL'),
    path: z.string()
        .min(1, 'Path is required')
        .regex(/^\/.*/, 'Path must start with /'),
    enabled: z.boolean(),
});

type RouteFormData = z.infer<typeof routeSchema>;

interface RouteFormProps {
    onSubmit?: (data: RouteFormData) => void;
    initialData?: Partial<RouteFormData>;
    isLoading?: boolean;
    onCancel?: () => void;
}

export default function RouteForm({
    onSubmit,
    initialData,
    isLoading = false,
    onCancel
}: RouteFormProps) {
    const [submitStatus, setSubmitStatus] = useState<'idle' | 'success' | 'error'>('idle');

    const {
        register,
        handleSubmit,
        formState: { errors, isValid, isDirty },
        reset,
        watch,
    } = useForm<RouteFormData>({
        resolver: zodResolver(routeSchema),
        defaultValues: {
            name: initialData?.name || '',
            host: initialData?.host || '',
            backend: initialData?.backend || '',
            path: initialData?.path || '/',
            enabled: initialData?.enabled || false,
        },
        mode: 'onChange',
    });

    const handleFormSubmit = async (data: RouteFormData) => {
        try {
            setSubmitStatus('idle');
            console.log('Route form data:', data);

            if (onSubmit) {
                await onSubmit(data);
            }

            setSubmitStatus('success');

            // Reset form if this is a create operation (no initial data)
            if (!initialData) {
                reset();
            }
        } catch (error) {
            console.error('Form submission error:', error);
            setSubmitStatus('error');
        }
    };

    const handleReset = () => {
        reset();
        setSubmitStatus('idle');
    };

    return (
        <Paper elevation={2} sx={{ p: 3, maxWidth: 600, mx: 'auto' }}>
            <Typography variant="h5" component="h2" gutterBottom>
                Route Configuration
            </Typography>

            {submitStatus === 'success' && (
                <Alert severity="success" sx={{ mb: 2 }}>
                    Route saved successfully!
                </Alert>
            )}

            {submitStatus === 'error' && (
                <Alert severity="error" sx={{ mb: 2 }}>
                    Failed to save route. Please try again.
                </Alert>
            )}

            <Box component="form" onSubmit={handleSubmit(handleFormSubmit)} noValidate>
                <TextField
                    {...register('name')}
                    label="Route Name"
                    fullWidth
                    margin="normal"
                    error={!!errors.name}
                    helperText={errors.name?.message}
                    placeholder="e.g., My API Route"
                />

                <TextField
                    {...register('host')}
                    label="Host"
                    fullWidth
                    margin="normal"
                    error={!!errors.host}
                    helperText={errors.host?.message || 'Domain where this route will be accessible'}
                    placeholder="e.g., api.example.com"
                />

                <TextField
                    {...register('backend')}
                    label="Backend URL"
                    fullWidth
                    margin="normal"
                    error={!!errors.backend}
                    helperText={errors.backend?.message || 'Full URL to the backend service'}
                    placeholder="e.g., https://backend.example.com"
                />

                <TextField
                    {...register('path')}
                    label="Path"
                    fullWidth
                    margin="normal"
                    error={!!errors.path}
                    helperText={errors.path?.message || 'Path pattern for routing'}
                    placeholder="e.g., /api/v1/users"
                />

                <Box sx={{ mt: 2, mb: 3 }}>
                    <FormControlLabel
                        control={
                            <Switch
                                {...register('enabled')}
                                color="primary"
                            />
                        }
                        label={`Enable Route`}
                    />
                </Box>

                <Box sx={{ display: 'flex', gap: 2, justifyContent: 'flex-end' }}>
                    {onCancel && (
                        <Button
                            type="button"
                            variant="outlined"
                            onClick={onCancel}
                            disabled={isLoading}
                        >
                            Cancel
                        </Button>
                    )}
                    <Button
                        type="button"
                        variant="outlined"
                        onClick={handleReset}
                        disabled={!isDirty || isLoading}
                    >
                        Reset
                    </Button>
                    <Button
                        type="submit"
                        variant="contained"
                        disabled={!isValid || isLoading}
                        sx={{ minWidth: 120 }}
                    >
                        {isLoading ? 'Saving...' : initialData ? 'Update Route' : 'Create Route'}
                    </Button>
                </Box>
            </Box>
        </Paper>
    );
}