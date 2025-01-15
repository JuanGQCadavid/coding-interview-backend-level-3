
/*

     This code was generated using ChatGPT
     The task was to change from javsscript server to the golang server.

     Original file is located in .back folder

*/


import fetch from 'node-fetch'; // Ensure you install node-fetch: npm install node-fetch
describe('E2E Tests', () => {
    const serverAddress = 'http://127.0.0.1:8000'; // Replace with your Go server's address and port

    type Item = {
        id: number
        name: string
        price: number
    };

    it('should get a response with status code 200', async () => {
        const response = await fetch(`${serverAddress}/ping`, {
            method: 'GET',
        });
        const result = await response.json();

        expect(response.status).toBe(200);
        expect(result).toEqual({ ok: true });
    });

    describe("Basic Items functionality", () => {
        it("should be able to list all items", async () => {
            const response = await fetch(`${serverAddress}/items`, {
                method: 'GET',
            });
            const result = await response.json();

            expect(response.status).toBe(200);
            expect(result).toEqual([]);

            await fetch(`${serverAddress}/items`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    name: 'Item 1',
                    price: 10
                }),
            });

            const response2 = await fetch(`${serverAddress}/items`, {
                method: 'GET',
            });
            const result2 = await response2.json();

            expect(response2.status).toBe(200);
            expect(result2).toEqual([{
                id: expect.any(Number),
                name: 'Item 1',
                price: 10
            }]);
        });

        it("should be able to create a new item and get it by id", async () => {
            const response = await fetch(`${serverAddress}/items`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    name: 'Item 1',
                    price: 10
                }),
            });
            const result = await response.json();

            expect(response.status).toBe(201);
            expect(result).toEqual({
                id: expect.any(Number),
                name: 'Item 1',
                price: 10
            });

            const response2 = await fetch(`${serverAddress}/items/${result.id}`, {
                method: 'GET',
            });
            const result2 = await response2.json();

            expect(response2.status).toBe(200);
            expect(result2).toEqual({
                id: result.id,
                name: 'Item 1',
                price: 10
            });
        });

        it("should be able to update an item", async () => {
            const response = await fetch(`${serverAddress}/items`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    name: 'Item 1',
                    price: 10
                }),
            });
            const createdItem = await response.json();

            const updateResponse = await fetch(`${serverAddress}/items/${createdItem.id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    name: 'Item 1 updated',
                    price: 20
                }),
            });
            const updatedItem = await updateResponse.json();

            expect(updateResponse.status).toBe(200);
            expect(updatedItem).toEqual({
                id: createdItem.id,
                name: 'Item 1 updated',
                price: 20
            });

            const response2 = await fetch(`${serverAddress}/items/${createdItem.id}`, {
                method: 'GET',
            });
            const result2 = await response2.json();

            expect(response2.status).toBe(200);
            expect(result2).toEqual(updatedItem);
        });

        it("should be able to delete an item", async () => {
            const response = await fetch(`${serverAddress}/items`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    name: 'Item 1',
                    price: 10
                }),
            });
            const createdItem = await response.json();

            const deleteResponse = await fetch(`${serverAddress}/items/${createdItem.id}`, {
                method: 'DELETE',
            });

            expect(deleteResponse.status).toBe(204);

            const response2 = await fetch(`${serverAddress}/items/${createdItem.id}`, {
                method: 'GET',
            });

            expect(response2.status).toBe(404);
        });
    });

    describe("Validations", () => {
        it("should validate required fields", async () => {
            const response = await fetch(`${serverAddress}/items`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    name: 'Item 1',
                }),
            });
            const result = await response.json();

            expect(response.status).toBe(400);
            expect(result).toEqual({
                errors: [
                    {
                        field: 'price',
                        message: 'Field "price" is required'
                    }
                ]
            });
        });

        it("should not allow negative pricing for new items", async () => {
            const response = await fetch(`${serverAddress}/items`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    name: 'Item 1',
                    price: -10
                }),
            });
            const result = await response.json();

            expect(response.status).toBe(400);
            expect(result).toEqual({
                errors: [
                    {
                        field: 'price',
                        message: 'Field "price" cannot be negative'
                    }
                ]
            });
        });

        it("should not allow negative pricing for updated items", async () => {
            const response = await fetch(`${serverAddress}/items`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    name: 'Item 1',
                    price: 10
                }),
            });
            const createdItem = await response.json();

            const updateResponse = await fetch(`${serverAddress}/items/${createdItem.id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    name: 'Item 1 updated',
                    price: -20
                }),
            });
            const result = await updateResponse.json();

            expect(updateResponse.status).toBe(400);
            expect(result).toEqual({
                errors: [
                    {
                        field: 'price',
                        message: 'Field "price" cannot be negative'
                    }
                ]
            });
        });
    });
});