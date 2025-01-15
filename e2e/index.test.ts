
/*

     This code was generated using ChatGPT
     The task was to change from javsscript server to the golang server.

     Original file is located in .back folder

*/
import fetch from 'node-fetch'

type Item = {
    id: number;
    name: string;
    price: number;
};

describe('E2E Tests', () => {
    const serverAddress = 'http://127.0.0.1:8000'; // Replace with your Go server's address and port

    it('should get a response with status code 200', async () => {
        const response = await fetch(`${serverAddress}/ping`, {
            method: 'GET',
        });
        const result = await response.json();

        expect(response.status).toBe(200);
        expect(result).toEqual({ ok: true });
    });

    describe('Basic Items functionality', () => {
        it('should be able to list all items', async () => {
            let response = await fetch(`${serverAddress}/items`, {
                method: 'GET',
            });
            let items: Item[] = await response.json();

            expect(response.status).toBe(200);
            expect(items).toEqual([]);

            response = await fetch(`${serverAddress}/items`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name: 'Item 1', price: 10 }),
            });
            const createdItem: Item = await response.json();

            response = await fetch(`${serverAddress}/items`, { method: 'GET' });
            items = await response.json();

            expect(response.status).toBe(200);
            expect(items).toContainEqual(createdItem);
        });

        it('should be able to create a new item and get it by id', async () => {
            const createResponse = await fetch(`${serverAddress}/items`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name: 'Item 1', price: 10 }),
            });
            const createdItem: Item = await createResponse.json();

            expect(createResponse.status).toBe(201);
            expect(createdItem).toEqual({
                id: expect.any(String),
                name: 'Item 1',
                price: 10,
            });

            const getResponse = await fetch(`${serverAddress}/items/${createdItem.id}`, {
                method: 'GET',
            });
            const fetchedItem: Item = await getResponse.json();

            expect(getResponse.status).toBe(200);
            expect(fetchedItem).toEqual(createdItem);
        });

        it('should be able to update an item', async () => {
            const createResponse = await fetch(`${serverAddress}/items`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name: 'Item 1', price: 10 }),
            });
            const createdItem: Item = await createResponse.json();

            const updateResponse = await fetch(`${serverAddress}/items/${createdItem.id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name: 'Updated Item', price: 20 }),
            });
            const updatedItem: Item = await updateResponse.json();

            expect(updateResponse.status).toBe(200);
            expect(updatedItem).toEqual({
                id: createdItem.id,
                name: 'Updated Item',
                price: 20,
            });
        });

        it('should be able to delete an item', async () => {
            const createResponse = await fetch(`${serverAddress}/items`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name: 'Item 1', price: 10 }),
            });
            const createdItem: Item = await createResponse.json();

            const deleteResponse = await fetch(`${serverAddress}/items/${createdItem.id}`, {
                method: 'DELETE',
            });

            expect(deleteResponse.status).toBe(204);

            const getResponse = await fetch(`${serverAddress}/items/${createdItem.id}`, {
                method: 'GET',
            });

            expect(getResponse.status).toBe(404);
        });
    });

    describe('Validations', () => {
        it('should validate required fields', async () => {
            const response = await fetch(`${serverAddress}/items`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name: 'Item 1' }), // Missing price
            });
            const result = await response.json();

            expect(response.status).toBe(400);
            expect(result).toEqual({
                errors: [{ field: 'price', message: 'Field "price" is required' }],
            });
        });

        it('should not allow negative pricing for new items', async () => {
            const response = await fetch(`${serverAddress}/items`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name: 'Item 1', price: -10 }),
            });
            const result = await response.json();

            expect(response.status).toBe(400);
            expect(result).toEqual({
                errors: [{ field: 'price', message: 'Field "price" cannot be negative' }],
            });
        });

        it('should not allow negative pricing for updated items', async () => {
            const createResponse = await fetch(`${serverAddress}/items`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name: 'Item 1', price: 10 }),
            });
            const createdItem: Item = await createResponse.json();

            const response = await fetch(`${serverAddress}/items/${createdItem.id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name: 'Updated Item', price: -20 }),
            });
            const result = await response.json();

            expect(response.status).toBe(400);
            expect(result).toEqual({
                errors: [{ field: 'price', message: 'Field "price" cannot be negative' }],
            });
        });
    });
});